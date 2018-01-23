package twelve

const testVersion = 1

var days = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var gifts = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

func Song() string {
	var song string
	for day := range days {
		song += Verse(day+1) + "\n"
	}
	return song
}

func Verse(day int) string {
	var verse string
	day--
	verse += "On the " + days[day] + " day of Christmas my true love gave to me, "
	if day == 0 {
		verse += gifts[day] + "."
	} else {
		for gift := range gifts[:day+1] {
			verse += gifts[day-gift]
			if gift == day-1 {
				verse += " and "
			} else if gift == day {
				continue
			} else {
				verse += ", "
			}
		}
		verse += "."
	}
	return verse
}
