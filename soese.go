///////////////////////////////////////////////////////////////////////////////////////
//              soese · stephen's over-engineered scheme engine · v0.0.0             //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////
//                                                                                   //
//                                      old code                                     //
//                                                                                   //
///////////////////////////////////////////////////////////////////////////////////////

// =================================================================
// DATA STRUCTURES & TYPES
// =================================================================

// We use type aliases for clarity.
type Symbol string
type Number int64
type List []any

// Proc is a built-in procedure implemented in Go.
// It receives its arguments already evaluated.
type Proc func(args List) (any, error)

// SpecialForm is a built-in procedure that does NOT have its arguments
// evaluated before being called. It's responsible for its own evaluation logic.
// Examples: if, define, lambda.
type SpecialForm func(args List, env *Env) (any, error)

// A Procedure represents a user-defined lambda. It captures the
// parameters, the function body, and the environment (closure)
// in which it was created.
type Procedure struct {
	params []string
	body   any
	env    *Env
}

// Env represents our environment, mapping symbols (strings) to values.
// It includes a reference to an outer environment to allow for lexical scoping.
type Env struct {
	store map[string]any
	outer *Env
}

// =================================================================
// ENVIRONMENT
// =================================================================

// NewEnv creates a new environment, optionally extending an outer one.
func NewEnv(outer *Env) *Env {
	return &Env{
		store: make(map[string]any),
		outer: outer,
	}
}

// Find retrieves the environment where a given variable is defined.
// It searches from the current environment outwards.
func (e *Env) Find(key string) *Env {
	if _, ok := e.store[key]; ok {
		return e
	}
	if e.outer != nil {
		return e.outer.Find(key)
	}
	return nil
}

// Get retrieves a value for a key from the environment chain.
func (e *Env) Get(key string) (any, bool) {
	env := e.Find(key)
	if env == nil {
		return nil, false
	}
	val, ok := env.store[key]
	return val, ok
}

// Set adds a new key-value pair to the current environment's store.
func (e *Env) Set(key string, val any) {
	e.store[key] = val
}

// =================================================================
// PARSER
// =================================================================

// tokenize converts an input string into a slice of tokens.
func tokenize(input string) []string {
	r := strings.NewReplacer("(", " ( ", ")", " ) ")
	return strings.Fields(r.Replace(input))
}

// parseAtom attempts to convert a single token into an atom (Number or Symbol).
func parseAtom(token string) any {
	if val, err := strconv.ParseInt(token, 10, 64); err == nil {
		return Number(val)
	}
	return Symbol(token)
}

// readFromTokens recursively builds an abstract syntax tree (AST).
func readFromTokens(tokens []string) (any, []string, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("syntax error: unexpected EOF")
	}
	token := tokens[0]
	tokens = tokens[1:]

	switch token {
	case "(":
		var list List
		for len(tokens) > 0 && tokens[0] != ")" {
			expr, rest, err := readFromTokens(tokens)
			if err != nil {
				return nil, nil, err
			}
			list = append(list, expr)
			tokens = rest
		}
		if len(tokens) == 0 || tokens[0] != ")" {
			return nil, nil, fmt.Errorf("syntax error: missing ')'")
		}
		return list, tokens[1:], nil
	case ")":
		return nil, nil, fmt.Errorf("syntax error: unexpected ')'")
	default:
		return parseAtom(token), tokens, nil
	}
}

// Parse orchestrates the tokenizing and parsing process.
func Parse(input string) (any, error) {
	tokens := tokenize(input)
	ast, remaining, err := readFromTokens(tokens)
	if err != nil {
		return nil, err
	}
	if len(remaining) > 0 {
		return nil, fmt.Errorf("syntax error: unexpected tokens at end of input: %v", remaining)
	}
	return ast, nil
}

// =================================================================
// EVALUATOR
// =================================================================

// Eval is the core of the interpreter. It evaluates an expression in a given environment.
// It is now fully recursive, replacing the previous tail-call-optimized loop.
func Eval(exp any, env *Env) (any, error) {
	switch e := exp.(type) {
	case Symbol:
		// Look up the symbol in the environment.
		val, ok := env.Get(string(e))
		if !ok {
			return nil, fmt.Errorf("error: undefined symbol '%s'", e)
		}
		return val, nil
	case Number:
		// Numbers evaluate to themselves.
		return e, nil
	case List:
		// An empty list is an error.
		if len(e) == 0 {
			return nil, fmt.Errorf("error: empty list cannot be evaluated")
		}

		first := e[0]
		args := e[1:]

		// Check if the expression is a call to a special form.
		// Special forms are looked up by symbol but are not evaluated like regular procedures.
		if sym, ok := first.(Symbol); ok {
			if proc, ok := env.Get(string(sym)); ok {
				if sf, ok := proc.(SpecialForm); ok {
					return sf(args, env)
				}
			}
		}

		// It's a regular procedure call. First, evaluate the procedure itself.
		proc, err := Eval(first, env)
		if err != nil {
			return nil, err
		}

		// Then, evaluate all arguments.
		evaledArgs := List{}
		for _, arg := range args {
			val, err := Eval(arg, env)
			if err != nil {
				return nil, err
			}
			evaledArgs = append(evaledArgs, val)
		}

		// Finally, apply the procedure to the evaluated arguments.
		return apply(proc, evaledArgs)
	default:
		return nil, fmt.Errorf("error: unknown expression type: %T", exp)
	}
}

// apply handles the logic of calling a procedure (built-in or user-defined).
func apply(proc any, args List) (any, error) {
	switch p := proc.(type) {
	case Proc:
		// It's a built-in Go function.
		return p(args)
	case Procedure:
		// It's a user-defined procedure.
		// Create a new environment for the function call, linked to the procedure's closure.
		localEnv := NewEnv(p.env)
		if len(p.params) != len(args) {
			return nil, fmt.Errorf("error: expected %d arguments, but got %d", len(p.params), len(args))
		}
		// Bind arguments to parameters in the new environment.
		for i, param := range p.params {
			localEnv.Set(param, args[i])
		}
		// Evaluate the procedure's body in the new, extended environment.
		return Eval(p.body, localEnv)
	default:
		return nil, fmt.Errorf("error: not a procedure: %v", proc)
	}
}

// =================================================================
// BUILT-IN PROCEDURES & GLOBAL ENVIRONMENT
// =================================================================

// createGlobalEnv initializes the top-level environment.
func createGlobalEnv() *Env {
	env := NewEnv(nil)

	// --- Special Forms ---

	env.Set("if", SpecialForm(func(args List, env *Env) (any, error) {
		if len(args) < 2 || len(args) > 3 {
			return nil, fmt.Errorf("syntax error: 'if' requires 2 or 3 arguments")
		}
		testResult, err := Eval(args[0], env)
		if err != nil {
			return nil, err
		}

		// In Scheme, only #f is false. Here, we'll treat a Go bool `false` as false.
		isFalse := false
		if b, ok := testResult.(bool); ok && !b {
			isFalse = true
		}

		if isFalse {
			if len(args) == 3 {
				return Eval(args[2], env) // Eval else-expression
			}
			return nil, nil // No else branch, return nil-like value
		}
		return Eval(args[1], env) // Eval then-expression
	}))

	env.Set("define", SpecialForm(func(args List, env *Env) (any, error) {
		if len(args) < 2 {
			return nil, fmt.Errorf("syntax error: 'define' requires at least 2 arguments")
		}

		// Handle function shorthand: (define (f x) body)
		if def, ok := args[0].(List); ok {
			if len(def) == 0 {
				return nil, fmt.Errorf("syntax error: invalid function definition")
			}
			funcName, ok := def[0].(Symbol)
			if !ok {
				return nil, fmt.Errorf("syntax error: function name must be a symbol")
			}
			// Desugar into (define f (lambda (params...) body))
			lambdaExp := List{Symbol("lambda"), def[1:], args[1]}
			proc, err := Eval(lambdaExp, env) // Eval the created lambda
			if err != nil {
				return nil, err
			}
			env.Set(string(funcName), proc)
			return nil, nil // 'define' returns an unspecified value
		}

		// Handle variable definition: (define var val)
		sym, ok := args[0].(Symbol)
		if !ok {
			return nil, fmt.Errorf("syntax error: 'define' key must be a symbol")
		}
		if len(args) != 2 {
			return nil, fmt.Errorf("syntax error: 'define' for a variable requires 2 arguments")
		}
		val, err := Eval(args[1], env)
		if err != nil {
			return nil, err
		}
		env.Set(string(sym), val)
		return nil, nil // 'define' returns an unspecified value
	}))

	env.Set("lambda", SpecialForm(func(args List, env *Env) (any, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("syntax error: lambda requires 2 arguments (params and body)")
		}
		paramsRaw, ok := args[0].(List)
		if !ok {
			return nil, fmt.Errorf("syntax error: lambda parameters must be a list")
		}
		var params []string
		for _, p := range paramsRaw {
			if ps, ok := p.(Symbol); ok {
				params = append(params, string(ps))
			} else {
				return nil, fmt.Errorf("syntax error: lambda parameters must be symbols")
			}
		}
		body := args[1]
		return Procedure{params: params, body: body, env: env}, nil
	}))

	// --- Regular Procedures ---

	env.Set("+", Proc(func(args List) (any, error) {
		sum := Number(0)
		for _, arg := range args {
			n, ok := arg.(Number)
			if !ok {
				return nil, fmt.Errorf("type error: '+' expects numbers, got %T", arg)
			}
			sum += n
		}
		return sum, nil
	}))

	env.Set("print", Proc(func(args List) (any, error) {
		var parts []string
		for _, arg := range args {
			parts = append(parts, stringify(arg))
		}
		fmt.Println(strings.Join(parts, " "))
		return nil, nil
	}))

	return env
}

// stringify converts an evaluated expression back into a readable string.
func stringify(exp any) string {
	if exp == nil {
		return ""
	}
	switch e := exp.(type) {
	case Number:
		return fmt.Sprintf("%d", e)
	case Symbol:
		return string(e)
	case List:
		var parts []string
		for _, item := range e {
			parts = append(parts, stringify(item))
		}
		return "(" + strings.Join(parts, " ") + ")"
	case Procedure:
		return "<procedure>"
	case SpecialForm:
		return "<special-form>"
	default:
		return fmt.Sprintf("%v", e)
	}
}

// =================================================================
// MAIN (REPL)
// =================================================================

func main() {
	globalEnv := createGlobalEnv()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Simple Scheme Interpreter in Go. Press Ctrl+D to exit.")

	for {
		fmt.Print("λ> ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil && err != io.EOF {
				fmt.Fprintln(os.Stderr, "error reading input:", err)
			}
			break
		}
		line := scanner.Text()
		if line == "" {
			continue
		}

		ast, err := Parse(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		result, err := Eval(ast, globalEnv)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		output := stringify(result)
		if output != "" {
			fmt.Println(output)
		}
	}
}
