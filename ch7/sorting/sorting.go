package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Demon Cleaner", "Kyuss", "Welcome To Sky Valley", 1994, length("5m19s")},
	{"Lost In The Cold Sun", "Earthless", "Sonic Prayer", 2005, length("20m25s")},
	{"Jam No. Phi", "The Machine", "Drie", 2010, length("10m23s")},
	{"The Flood", "Mustard Gas And Roses", "Becoming", 2016, length("4m35s")},
	{"March To The Sea", "Pelican", "The Fire In Our Throat Will Beckon The Thaw", 2005, length("16m3s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "------", "------", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	printTracks(tracks)
	sort.Sort(byArtist(tracks))
	printTracks(tracks)
	sort.Sort(byYear(tracks))
	printTracks(tracks)
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		return false
	}})
	printTracks(tracks)
}
