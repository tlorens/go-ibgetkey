C binding for Go to read keystrokes one at a time and return it's ASCII value. 

This package is for command-line applications when you may want to ask for a passwords and echo
back asterisks or want a hotkey effect; Single key input without pressing Enter to continue.  Use
this in a for/while loop to build input strings, etc. etc.  This function works much like Turbo Pascal's
ReadKey() function does. 

## Usage

```
for {
	// Good for special keys: Arrows, PgUp, PgDn, Home, End.
	ch := keyboard.ReadKey()
	fmt.Printf("[%d][%c]\n", ch, ch);	
}
```

```
for {
	// Use this if you want results untoched
	// You will have to process/parse special keys on your own.
	ch := keyboard.RawKey()
	fmt.Printf("[%d][%c]\n", ch, ch);	
}
	
```


I've put this out here to share.  Do what you like.  Maybe someone is/was
having the same problem with reading keycodes as I was. 