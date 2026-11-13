package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = true

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var text2 = `Известный физик и автор книг-бестселлеров Лоуренс Краусс 
	предлагает новаторский взгляд на то, как возникло в начале времен все, 
	что существует вокруг нас. Откуда взялась Вселенная? Что было до нее? 
	Что принесет будущее? И наконец, почему вообще существует что-то, а не ничто? 
	Краусс – один из немногих видных ученых наших дней, кому удалось преодолеть 
	пропасть между наукой и популярной культурой. Он описывает поразительно 
	красивые экспериментальные наблюдения и головоломные новые теории, 
	которые не только наглядно показывают, что нечто может возникнуть из ничего, 
	но и утверждают, что нечто возникает из ничего всегда. Книга «Всё из ничего: 
	Как возникла Вселенная», снабженная новым предисловием о том, какое значение 
	имеет открытие бозона Хиггса, отличается характерным для Краусса мрачным 
	юмором и необыкновенно ясными объяснениями, которые переносят нас к началу 
	начал и представляют новейшие свидетельства того, как развивалась наша Вселенная, 
	– и указания на то, чем она, скорее всего, кончит.`

var englishText = `The renowned physicist 🔬 and bestselling author Lawrence Krauss 
	offers 🔬 a groundbreaking view 🔬 of how everything around us came into existence 
	at the beginning of time ⏳. Where did the universe come from? 🌌 What was there 
	before it? ❓ What will the future bring? 🔮 And finally, why is there something 
	rather than nothing at all? 🤔 Krauss is one of the few prominent scientists of 
	our day to have bridged the gap between science and popular culture 🌍. He 
	describes stunningly beautiful experimental observations ✨ and mind-bending 
	new theories 🧠 that not only demonstrate clearly that something can arise 
	from nothing 🕳️➡️💥, but also assert that something always arises from nothing ♾️. 
	The book 📖 📖 📖 "A Universe from Nothing: Why There Is Something Rather Than Nothing" 📖 📖, 
	equipped with a new preface 📖 📖 on the significance of the discovery of the 
	Higgs boson ⚛️, is distinguished by Krauss's characteristic dark humor 😏 and 
	extraordinarily clear explanations 💡 that carry us back to the very beginning 
	of beginnings 🌀 and present the latest evidence of how our universe evolved 🌠—and 
	indications of how it will most likely end 💫.`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})
}

func TestTop5(t *testing.T) {
	t.Run("no words", func(t *testing.T) {
		require.Len(t, Top5(""), 0)
	})

	t.Run("correct text", func(t *testing.T) {
		expected := []string{
			"и",         // 6
			"что",       // 6
			"из",        // 5
			"вселенная", // 3
			"как",       // 3
		}
		require.Equal(t, expected, Top5(text2))
	})
}

func TestTop10English(t *testing.T) {
	t.Run("no words", func(t *testing.T) {
		require.Len(t, Top10Eng(""), 0)
	})

	t.Run("correct test", func(t *testing.T) {
		expected := []string{
			"the",
			"of",
			"📖",
			"and",
			"nothing",
			"from",
			"is",
			"something",
			"that",
			"a",
		}
		require.Equal(t, expected, Top10Eng(englishText))
	})
}
