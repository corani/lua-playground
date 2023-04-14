# lua-playground
Type some Lua code in the input box and press "Submit" to send it to the server. There it's executed
using [yuin/gopher-lua](https://github.com/yuin/gopher-lua) and the console output is returned.

NOTE: This is just a toy project to play around with `gopher-lua`.

## Usage
```console
docker compose up
```

NOTE: If you need to use a `GOPROXY` to fetch the dependencies, create a `.env` file containing
`GOPROXY=<address of proxy>`.

## Demo
![image](https://user-images.githubusercontent.com/480775/231712301-6f873a50-10bb-40fc-bfda-53ac39593fd5.png)
