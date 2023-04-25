# Monty Hall

## Context

The Monty Hall problem is a probabilistic puzzle loosely based on the TV game show "Let's Make a Deal". It is named after the man who introduced the game, Monty Hall. The game goes like this: In each game, three doors are presented to the player, behind one of the three doors there is a prize and behind the other two doors there is nothing. Monty Hall then asks the player to choose a door. The rule of the game is that the player will be given the prize hiding behind the door they choose. After the player chooses a door, say *Door 1*, Monty Hall takes a look at what is behind the other two doors and opens one of them. He says, "Look, there are no prizes behind this door!". Since only one door has a prize, it is always possible for Monty Hall to choose a door other than the player's door and that has no prize. Let's assume that the door he opens is *Door 2*. Monty Hall then goes on to say, I will make you an offer: I will allow you to change your decision by changing your choice to the other door (in this example, to *Door 3*). Do you accept the offer? It's up to the player to decide whether to switch to Gate 3 or stay with *Gate 1*. After the player decides to "switch gates or not", Monty Hall opens the final choice gate. The player will get a prize if and only if there is a prize behind the door of his choice.

Historically there have been intense debates about whether changing the initial choice would improve the chance of winning, which has made the television show very popular. In this lab, you must develop a program that simulates the Monty Hall game and examine this problem.

## Description

The program you have to write has two modes: the game mode and the search mode. Briefly, in the game mode, the program will act as Monty Hall and play the game with the program user. In the search mode, the program allows the user to study what is the probability of winning the prize for the strategy of systematically changing the initial choice and without making any changes.

When the program runs, it first displays a menu and asks the user to select an option:

Enter one of the following options:

1. Game mode

2. Search mode

3. Output

Only when the user chooses to exit will the program exit. If the user chooses game mode and a game ends, the program returns to the menu and prompts the user to enter an option again. Similarly if the user chooses the search mode and the search for either strategy ends, the program also returns to the menu and asks the user again.

### Game mode

In the game mode, the program acts as Monty Hall and simulates a game with the user. The three doors will be denoted as Door 1, Door 2 and Door 3. The program begins by selecting one door from the three in a (uniformly) random fashion by putting the prize behind the resulting door. It then asks the user to choose a door. After the user makes his choice, the program, based on the user's choices, "opens" a door that the user did not choose and that does not hide a prize. Then, the program asks the user if he wants to change. After the user decides to switch or not, the program reveals where the prize is and announces whether the user has won. In addition to this main task, the total number of games the user has played in the game mode and the number of times he wins should also be counted and printed after each game. Note that even if a user exits game mode to research mode, their statistics in game mode (i.e., the total number of games played and the number of wins in game mode so far) should still be retained. When the user returns to the game mode, the account can be continued.

### Search mode

In the search mode, the program first asks the user to choose one of two strategies: systematically changing the initial choice and making no change. It also asks the user to enter the number (N) of games that the user would like to simulate for this strategy. Then, for the chosen strategy, the program simulates N games that are played using this strategy. In each game, the program again acts like Monty Hall except that

1. No message is printed;

2. The program always sets the player's initial choice to Gate 1;

3. If the *"never change"* strategy is simulated, the program sets the player's choice of "*change or no change"* always to *"no change"*. If the strategy *"always change"* is simulated, the program sets the player's choice *"change or no change"* always to "*change"*.

In the search mode, the actions of the player are also simulated by the program and it is not necessary for the user to enter anything in a game. After the simulation of N games, the program prints the percentage where the candidate wins (which approximates the probability of winning, for large N).

### Requirements

In addition to the regular requirements for labs, this lab requires you to modularize your code in terms of functions. In particular, note that game mode and search mode have some things in common. You need to extract the common parts and put them into functions and then call them in both modes of the program.