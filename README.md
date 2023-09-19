# Simple WASM - GO WebAssembly
<img src="https://img.shields.io/badge/Type-Web%2FSitus%20web-lightgrey" alt="Type : Web/Situs web"> 

## Halo👋

Language: English

## Description:
  Just a little try to create simple WebAssembly using GO. 

## Structure
  So i'll explain my understanding about this project, here's the structure how WebAssembly GO Works:<br><br>
  
  >__Server__\
  path: ```server/main.go``` this is for the web server, wich will catch **address** and throw back some WebAssembly Component that will be display to user (In this implementation i use ```8081``` port). Just run this, and the server will be active.

  >__WebAssembly GO__\
  path: ```wasm/main.go``` this is for the WebAssembly code, wich we can build to ```.wasm``` file using:<br> ```GOOS=js GOARCH=wasm go build -o ../main.wasm```.

  All of these components is interrelated,<br>
  1. ```Client/User``` **--(sending request address)-->** ```server/main.go``` <br>
  Then server will preparing the WASM,<br>
  ```server/main.go``` **--(preparing)-->** ```WASM Component``` (index.html, script.js, main.wasm)

  2. WASM Preparing process: <br>
  ```index.html``` **--(get WASM action through js)-->** ```script.js``` **--(js Get WASM action)-->** ```main.wasm```

  3. And finally, <br>
  ```main.wasm``` will return some action to ```index.html``` through ```script.js``` so that client/user can receive the requested action in index.html

<br>**Youtube Reference**: [Go and WebAssembly | Learn the Basics of WASM](https://www.youtube.com/watch?v=10Mz3z-W1BE)

## Instruction : 
- Clone this repo ( ```git clone https://github.com/Khip01/Simple-WASM_Go.git``` )
- Then run the server/main.go

## Code Editor :
- Visual Studio Code
 
# Use
[![My Skills](https://skillicons.dev/icons?i=go,vscode,wasm)](https://github.com/Khip01)
