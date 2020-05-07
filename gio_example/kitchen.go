// SPDX-License-Identifier: Unlicense OR MIT

package main

// A Gio program that demonstrates Gio widgets. See https://gioui.org for more information.

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/app/headless"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"golang.org/x/exp/shiny/materialdesign/icons"
)

var screenshot = flag.String("screenshot", "", "save a screenshot to a file and exit")

type scaledConfig struct {
	Scale float32
}

type iconAndTextButton struct {
	theme *material.Theme
}

func main() {
	flag.Parse()
	editor.SetText(longText)
	ic, err := widget.NewIcon(icons.ContentAdd)
	if err != nil {
		log.Fatal(err)
	}
	icon = ic
	progressIncrementer = make(chan int)
	gofont.Register()
	if *screenshot != "" {
		if err := saveScreenshot(*screenshot); err != nil {
			fmt.Fprintf(os.Stderr, "failed to save screenshot: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	go func() {
		for {
			time.Sleep(time.Second)
			progressIncrementer <- 10
		}
	}()

	go func() {
		w := app.NewWindow(app.Size(unit.Dp(800), unit.Dp(650)))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func saveScreenshot(f string) error {
	const scale = 1.5
	sz := image.Point{X: 800 * scale, Y: 600 * scale}
	w, err := headless.NewWindow(sz.X, sz.Y)
	if err != nil {
		return err
	}
	gtx := new(layout.Context)
	gtx.Reset(nil, &scaledConfig{scale}, sz)
	th := material.NewTheme()
	kitchen(gtx, th)
	w.Frame(gtx.Ops)
	img, err := w.Screenshot()
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return err
	}
	return ioutil.WriteFile(f, buf.Bytes(), 0666)
}

func loop(w *app.Window) error {
	th := material.NewTheme()

	gtx := new(layout.Context)
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx.Reset(e.Queue, e.Config, e.Size)
				kitchen(gtx, th)
				e.Frame(gtx.Ops)
			}
		case p := <-progressIncrementer:
			progress += p
			if progress > 100 {
				progress = 0
			}
			w.Invalidate()
		}
	}
}

var (
	editor     = new(widget.Editor)
	lineEditor = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	button            = new(widget.Button)
	greenButton       = new(widget.Button)
	iconTextButton    = new(widget.Button)
	iconButton        = new(widget.Button)
	radioButtonsGroup = new(widget.Enum)
	list              = &layout.List{
		Axis: layout.Vertical,
	}
	progress            = 0
	progressIncrementer chan int
	green               = true
	topLabel            = "Hello, Gio"
	icon                *widget.Icon
	checkbox            = new(widget.Bool)
	swtch               = new(widget.Bool)
)

func (b iconAndTextButton) Layout(gtx *layout.Context, button *widget.Button, icon *widget.Icon, word string) {
	material.ButtonLayout(b.theme).Layout(gtx, button, func() {
		iconAndLabel := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}
		textIconSpacer := unit.Dp(5)

		layIcon := layout.Rigid(func() {
			layout.Inset{Right: textIconSpacer}.Layout(gtx, func() {
				size := gtx.Px(unit.Dp(56)) - 2*gtx.Px(unit.Dp(16))
				if icon != nil {
					icon.Layout(gtx, unit.Px(float32(size)))
					gtx.Dimensions = layout.Dimensions{
						Size: image.Point{X: size, Y: size},
					}
				}
			})
		})

		layLabel := layout.Rigid(func() {
			layout.Inset{Left: textIconSpacer}.Layout(gtx, func() {
				widget.Label{}.Layout(gtx, b.theme.Shaper, text.Font{}, b.theme.TextSize, word)
			})
		})

		iconAndLabel.Layout(gtx, layIcon, layLabel)
	})
}

func kitchen(gtx *layout.Context, th *material.Theme) {
	widgets := []func(){
		func() {
			material.H3(th, topLabel).Layout(gtx)
		},
		func() {
			gtx.Constraints.Height.Max = gtx.Px(unit.Dp(200))
			material.Editor(th, "Hint").Layout(gtx, editor)
		},
		func() {
			e := material.Editor(th, "Hint")
			e.Font.Style = text.Italic
			e.Layout(gtx, lineEditor)
			for _, e := range lineEditor.Events(gtx) {
				if e, ok := e.(widget.SubmitEvent); ok {
					topLabel = e.Text
					lineEditor.SetText("")
				}
			}
		},
		func() {
			in := layout.UniformInset(unit.Dp(8))
			layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func() {
					in.Layout(gtx, func() {
						material.IconButton(th, icon).Layout(gtx, iconButton)
					})
				}),
				layout.Rigid(func() {
					in.Layout(gtx, func() {
						iconAndTextButton{th}.Layout(gtx, iconTextButton, icon, "Horizontal button")
					})
				}),
				layout.Rigid(func() {
					in.Layout(gtx, func() {
						for button.Clicked(gtx) {
							green = !green
						}
						material.Button(th, "Click me!").Layout(gtx, button)
					})
				}),
				layout.Rigid(func() {
					in.Layout(gtx, func() {
						btn := material.Button(th, "Green button")
						if green {
							btn.Background = color.RGBA{A: 0xff, R: 0x9e, G: 0x9d, B: 0x24}
						}
						btn.Layout(gtx, greenButton)
					})
				}),
			)
		},
		func() {
			material.ProgressBar(th).Layout(gtx, progress)
		},
		func() {
			layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func() {
					material.CheckBox(th, "Checkbox").Layout(gtx, checkbox)
				}),
				layout.Rigid(func() {
					layout.Inset{Left: unit.Dp(16)}.Layout(gtx, func() {
						material.Switch(th).Layout(gtx, swtch)
					})
				}),
			)
		},
		func() {
			layout.Flex{}.Layout(gtx,
				layout.Rigid(func() {
					material.RadioButton(th, "r1", "RadioButton1").Layout(gtx, radioButtonsGroup)
				}),
				layout.Rigid(func() {
					material.RadioButton(th, "r2", "RadioButton2").Layout(gtx, radioButtonsGroup)
				}),
				layout.Rigid(func() {
					material.RadioButton(th, "r3", "RadioButton3").Layout(gtx, radioButtonsGroup)
				}),
			)
		},
	}

	list.Layout(gtx, len(widgets), func(i int) {
		layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
	})
}

func (s *scaledConfig) Now() time.Time {
	return time.Now()
}

func (s *scaledConfig) Px(v unit.Value) int {
	scale := s.Scale
	if v.U == unit.UnitPx {
		scale = 1
	}
	return int(math.Round(float64(scale * v.V)))
}

const longText = `1. I learned from my grandfather, Verus, to use good manners, and to
put restraint on anger. 2. In the famous memory of my father I had a
pattern of modesty and manliness. 3. Of my mother I learned to be
pious and generous; to keep myself not only from evil deeds, but even
from evil thoughts; and to live with a simplicity which is far from
customary among the rich. 4. I owe it to my great-grandfather that I
did not attend public lectures and discussions, but had good and able
teachers at home; and I owe him also the knowledge that for things of
this nature a man should count no expense too great.

5. My tutor taught me not to favour either green or blue at the
chariot races, nor, in the contests of gladiators, to be a supporter
either of light or heavy armed. He taught me also to endure labour;
not to need many things; to serve myself without troubling others; not
to intermeddle in the affairs of others, and not easily to listen to
slanders against them.

6. Of Diognetus I had the lesson not to busy myself about vain things;
not to credit the great professions of such as pretend to work
wonders, or of sorcerers about their charms, and their expelling of
Demons and the like; not to keep quails (for fighting or divination),
nor to run after such things; to suffer freedom of speech in others,
and to apply myself heartily to philosophy. Him also I must thank for
my hearing first Bacchius, then Tandasis and Marcianus; that I wrote
dialogues in my youth, and took a liking to the philosopher's pallet
and skins, and to the other things which, by the Grecian discipline,
belong to that profession.

7. To Rusticus I owe my first apprehensions that my nature needed
reform and cure; and that I did not fall into the ambition of the
common Sophists, either by composing speculative writings or by
declaiming harangues of exhortation in public; further, that I never
strove to be admired by ostentation of great patience in an ascetic
life, or by display of activity and application; that I gave over the
study of rhetoric, poetry, and the graces of language; and that I did
not pace my house in my senatorial robes, or practise any similar
affectation. I observed also the simplicity of style in his letters,
particularly in that which he wrote to my mother from Sinuessa. I
learned from him to be easily appeased, and to be readily reconciled
with those who had displeased me or given cause of offence, so soon as
they inclined to make their peace; to read with care; not to rest
satisfied with a slight and superficial knowledge; nor quickly to
assent to great talkers. I have him to thank that I met with the
discourses of Epictetus, which he furnished me from his own library.

8. From Apollonius I learned true liberty, and tenacity of purpose; to
regard nothing else, even in the smallest degree, but reason always;
and always to remain unaltered in the agonies of pain, in the losses
of children, or in long diseases. He afforded me a living example of
how the same man can, upon occasion, be most yielding and most
inflexible. He was patient in exposition; and, as might well be seen,
esteemed his fine skill and ability in teaching others the principles
of philosophy as the least of his endowments. It was from him that I
learned how to receive from friends what are thought favours without
seeming humbled by the giver or insensible to the gift.`
