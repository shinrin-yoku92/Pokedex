# Pokédex 🧬

Welcome to **Pokedex**, a sleek command-line Pokédex built in **Go**.

---

## 🚀 Features
---

- **Catch Pokémon** from nearby map areas  
- **Inspect your Pokémon** with detailed info  
- **Map exploration** to discover new wild Pokémon  
- **Explore regions** using in-game navigation  
- **Built-in REPL** with intuitive commands  
- **Simple exit command** to end your session  

---

## 🧱 Project Structure
---

- **main.go** — program entrypoint  
- **repl.go** — main REPL loop  
- **command_*** — modular command handlers  
  - command_pokedex.go  
  - command_catch.go  
  - command_inspect.go  
  - command_map.go  
  - command_explore.go  
  - command_exit.go  
  - command_help.go  
- **repl_test.go** — REPL test suite  

---

## 🛠 Getting Started
---

### 1. Clone the repo  
git clone https://github.com/shinrin-yoku92/Pokedex.git  
cd Pokedex

### 2. Build  
go build -o pokedex

### 3. Run  
./pokedex

---

## 📚 Example Session
---

> help  
Available commands: catch, inspect, map, explore, exit

> map  
You look around the area… Pokémon spotted nearby!

> catch 25  
You caught a **Pikachu**!

> inspect pikachu  
Name: Pikachu  
Type: Electric  
Stats and more…

> exit  
Goodbye, Trainer!

---

## 🧪 Testing
---

Run all tests:

go test ./...

---

## 🔮 Possible Future Enhancements
---

- Pokémon evolution  
- Battle mechanics  
- Save/load Pokédex locally  
- Type effectiveness system  
- Larger region maps  

---

## 📜 License
---

Released under the **MIT License**.  
See the LICENSE file for full details.
