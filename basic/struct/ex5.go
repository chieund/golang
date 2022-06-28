package main

import "fmt"
type song struct {
	title, artist string
}

type playlist struct {
	genre string
	songs []song
}

func main()  {
	song1 := song{title: "wonderwall", artist: "oasis"}
	song2 := song{title: "super sonnic", artist: "oasis"}

	fmt.Printf("song1: %+v\nsong2: %+v\n", song1, song2)

	if song1 == song2 {
		fmt.Println("songs are equal.")
	} else {
		fmt.Println("songs are not equal.")
	}

	songs := [] song {
		{title: "wonderwall", artist: "oasis"},
		{title: "wonderwall1", artist: "oasis1"},
	}

	rock := playlist{
		genre: "indie rock",
		songs: songs,
	}

	song := rock.songs[0]
	song.title = "live forever"

	rock.songs[0].title = "live forever"
	fmt.Printf("\n%+v\n%+v\n", song, rock.songs[0])
	fmt.Printf("\n%-20s %20s\n", "TITLE", "ARTIST")
	for _, s:= range rock.songs {
		s.title = "destroy"
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}

	fmt.Printf("\n%-20s %20s\n", "TITLE", "ARTIST")
	for _, s:= range rock.songs {
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}
}
