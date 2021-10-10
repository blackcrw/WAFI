package internal

import "fmt"

func TextBanner() string {
	return fmt.Sprint(
		"   , ,, ,                                \n",
		"  | || |    ,/  _____  \\.               \n",
		"  \\_||_/    ||_/     \\_||              \n",
		"    ||       \\_| . . |_/                \n",
		"    ||         |  L  |                   \n",
		"   ,||         |`==='|                   \n",
		"   |>|      ___`>  -<'___                \n",
		"   |>|\\    /             \\             \n",
		"   \\>| \\  /  ,    .    .  |            \n",
		"    ||  \\/  /| .  |  . |  |             \n",
		"    ||\\  ` / | ___|___ |  |     (       \n",
		"  ((|| `--'  | _______ |  |     ))  (    \n",
		"( ) || (  )\\ | - --- - | —| (  ( \\  )) \n",
		"(\\/ || ))/ ( | -- - -- |  | )) )  \\((  \n",
		"( ()||((( ())|         |  |( (( () )     \n",
		"            — A K U M A —                \n",
		"           (Alpha Version)               \n")
}

func SimpleBanner() { fmt.Println(TextBanner()) }
