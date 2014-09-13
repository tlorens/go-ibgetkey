C binding for Go to read keystrokes one at a time and return it's ASCII value. 
Much like Turbo Pascals ReadKey() function. For use in a while loop.  


## Usage

```
ch := ibkey.ReadKey(false)
```

This is a bit hacky, but passing true will pass back keystrokes 
untouched (Raw mode).  While passing false will 'eat' ESC[ codes and 
return the keycode + 100.  


I've put this out here to share.  Do what you like.  Maybe someone is/was
having the same problem with reading keycodes as I was. 