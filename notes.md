# grid background

``` 
.visualGrid {
    position: fixed;
    width: calc(100vw - 20px);
    min-height: calc(100vh - 20px);
    top: 10px;
    left: 10px;
    background-image: radial-gradient(circle at 1px 1px, #000 1px, transparent 0px);
    background-size: calc((100vw - 20px) / 32) calc((100vw - 10px) / 32);
    pointer-events: none;
    z-index: -1;
    background-repeat: repeat;
}
```

this creates a dot:
```
background-image: radial-gradient(circle at 1px 1px, #000 1px, transparent 0px);
```

by making the size a small square (responsive to the viewport width), this dot gets repeated across the background since 
the default value for background-repeat is "repeat repeat", covering the whole page
```
background-size: calc((100vw - 20px) / 32) calc((100vw - 10px) / 32);
```

