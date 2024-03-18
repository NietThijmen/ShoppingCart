# Shopping Cart

## Description
A simple application allowing you to make a list of SSH hosts and connect to them using a single command.
This was meant as a simple project to manage all my SSH logins, and to learn about custom implementations of input & select dialogs in the terminal.

## Features
- Add, remove, and list SSH hosts
- Connect to a host using a single command
- Custom input and select dialogs

## Installation
1. Install the source code
```bash
git clone git@github.com:NietThijmen/ShoppingCart.git 
```

2. Compile the source code
```bash
cd ShoppingCart
go build
```

3. Make the program executable (linux & mac)
4. Add the following line to your `.bashrc` or `.zshrc` file:
```bash
alias ShoppingCart='/path/to/ShoppingCart'
```

## Usage
```bash
ShoppingCart [command]
```