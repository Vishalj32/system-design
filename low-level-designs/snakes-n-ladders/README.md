# Snake and Ladder Game

## Overview
This is a command-line-based Snake and Ladder game where players compete to be the first to reach the 100th cell of the board. The game is customizable and supports different configurations, such as the number of snakes, ladders, and players.

## How to Play

### Inputs:
The game accepts input in the following format:

1. **Number of Snakes (`s`)**: Followed by `s` lines each containing two integers representing the head and tail of the snake.
2. **Number of Ladders (`l`)**: Followed by `l` lines each containing two integers representing the start and end of the ladder.
3. **Number of Players (`p`)**: Followed by `p` lines each containing the name of a player.

### Example Input:
3 14 7 98 79 63 60 2 8 32 36 44 3 Alice Bob Charlie


In this example:
- 3 snakes: head at 14 (tail at 7), head at 98 (tail at 79), head at 63 (tail at 60).
- 2 ladders: start at 8 (end at 32), start at 36 (end at 44).
- 3 players: Alice, Bob, and Charlie.

### Game Rules:
1. Each player rolls a six-sided dice in turn.
2. Players move their pieces forward by the dice value.
3. If a player lands on the head of a snake, they move down to the snake’s tail.
4. If a player lands at the bottom of a ladder, they move up to the top of the ladder.
5. A player wins when they land exactly on the 100th cell.
6. If a player’s dice roll would move them beyond 100, they stay in their current position.

### Example Game Play:

Here’s an example of how the game might proceed:
<player_name> rolled a <dice_value> and moved from <initial_position> to <final_position>

When a player wins:
<player_name> wins the game!


## Optional Features
- Two dice gameplay for a roll range of 2-12.
- Customizable board size.
- Special rules for rolling 6.
- Continued play until only one player remains.