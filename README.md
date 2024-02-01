# fly-pokeapi

## Goal

The goal of this repository is to help you learn [Go](https://go.dev/) by building a CLI tool that interacts with the [Pokémon API](https://pokeapi.co/). With this tool, you should be able to execute the following commands:

`fly-pokeapi pokemon <name>`:  Name must be mandatory. This command gets the detail of a pokemon and prints a table in the terminal output containing the name of the Pokémon, its types, and its abilities.

*Example output:*

|   Name    |        Types         |                   Abilities                   |
|-----------|----------------------|-----------------------------------------------|
| Bulbasaur | Grass, Poison        | Overgrow, Chlorophyll                         |



`fly-pokeapi attack`: Use this command to group the first 9 Pokémon by type and sort them by their attack power. Emojis representing each type (🔥 🌵 💧) should be used.

*Example output:*

🔥 Fire:
|   Pokemon   | Attack |
|-------------|--------|
| Charizard   |   84   |
| Charmander  |   52   |
| Charmeleon  |   64   |

🌵 Grass:
|   Pokemon   | Attack |
|-------------|--------|
| Venusaur    |   82   |
| Bulbasaur   |   49   |
| Ivysaur     |   62   |

💧 Water:
|   Pokemon   | Attack |
|-------------|--------|
| Blastoise   |   83   |
| Squirtle    |   48   |
| Wartortle   |   63   |

`fly-pokeapi stats`: This command creates a CSV file containing the first 9 Pokémon and their evolutionary chain. Additionally, it groups and sums up the each stat belonging to each Pokémon.

```csv
Pokemon,Type,Health Points,Attack,Special Attack,Defense,Special Defense,Speed,Evolution
bulbasaur,Grass,45,49,65,49,65,45,ivysaur
ivysaur,Grass,60,62,80,63,80,60,venusaur
venusaur,Grass,80,82,100,83,100,80,None
charmander,Fire,39,52,60,43,50,65,charmeleon
charmeleon,Fire,58,64,80,58,65,80,charizard
charizard,Fire,78,84,109,78,85,100,None
squirtle,Fire,44,48,50,65,64,43,wartortle
wartortle,Fire,59,63,65,80,80,58,blastoise
blastoise,Fire,79,83,85,100,105,78,None
Aggregated,None,542,587,694,619,694,609,None
```


🌟🌟🌟 **Bonus**: Make sure to focus on making the output visually appealing and user-friendly.

This repository contains exercises designed to enhance your Go skills while engaging with the Pokémon API. Have fun exploring and building your CLI tool!