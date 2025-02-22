package wisdom

import (
	"context"
	"math/rand"
)

type InMemmoryRepoository struct {
	list []string
}

func (in *InMemmoryRepoository) Get(_ context.Context) (string, error) {
	n := rand.Intn(len(in.list))
	return in.list[n], nil
}

func NewInMemmoryRepository() *InMemmoryRepoository {
	return &InMemmoryRepoository{
		list: []string{
			"Science is organized knowledge. Wisdom is organized life. Immanuel Kant",
			"Doubt is the origin of wisdom. Augustine of Hippo",
			"The truest wisdom is a resolute determination. Napoleon Bonaparte",
			"Wisdom is not a product of schooling but of the lifelong attempt to acquire it. Albert Einstein",
			"Wisdom is the power to put our time and our knowledge to the proper use. Thomas J. Watson",
			"A symptom of wisdom is curiosity. The evidence is calmness and perseverance. The causes are experimentation and understanding. Maxime Lagacé",
			"It is not the man who has too little, but the man who craves more, that is poor. Seneca",
			"A wise man never loses anything, if he has himself. Michel de Montaigne",
			"A fool is known by his speech; and a wise man by silence. Pythagoras",
			"There is only a finger’s difference between a wise man and a fool. Diogenes",
			"Never say no twice if you mean it. Nassim Nicholas Taleb",
			"Irrigators channel waters; fletchers straighten arrows; carpenters bend wood; the wise master themselves. Buddha",
			"Wisdom is nothing but a preparation of the soul, a capacity, a secret art of thinking, feeling and breathing thoughts of unity at every moment of life. Hermann Hesse",
			"A lot of wisdom is just realizing the long-term consequences of your actions. The longer term you’re willing to look, the wiser you’re going to seem to everybody around you. Naval Ravikant",
			"Everything comes in time to him who knows how to wait. Leo Tolstoy",
			"Logic is the beginning of wisdom, not the end. Leonard Nimoy",
			"Knowledge speaks, but wisdom listens. Jimi Hendrix",
			"Silence is the sleep that nourishes wisdom. Francis Bacon",
			"A weak reaction is to rush things. A strong reaction is to go slow and steady. Maxime Lagacé",
			"The best words are the ones you are ready for. Maxime Lagacé",
			"We don’t receive wisdom; we must discover it for ourselves after a journey that no one can take for us or spare us. Marcel Proust",
			"Keep me away from the wisdom which does not cry, the philosophy which does not laugh and the greatness which does not bow before children. Kahlil Gibran",
			"A man has made at least a start on discovering the meaning of human life when he plants shade trees under which he knows full well he will never sit. David Elton Trueblood",
			"The wise know they are fools. Fools think they are wise. Maxime Lagacé",
			"words of wisdom discipline wisdom vice versa scott peck wisdom",
			"Discipline is wisdom and vice versa. M. Scott Peck",
			"The greatest wealth is to live content with little. Plato",
			"Wisdom is knowing when you can’t be wise. Paul Engle",
			"Wisdom consists of the anticipation of consequences. Norman Cousins",
			"Wisdom and deep intelligence require an honest appreciation of mystery. Thomas Moore",
			"The two powers which in my opinion constitute a wise man are those of bearing and forbearing. Epictetus",
			"Wisdom is keeping a sense of fallibility of all our views and opinions. Gerald Brenan",
			"Much wisdom often goes with fewest words. Sophocles",
			"Great wisdom is generous; petty wisdom is contentious. Zhuangzi",
			"Who then is free? The wise man who can command himself. Horace",
			"We don’t stop playing because we grow old, we grow old because we stop playing. George Bernard Shaw",
			"True wisdom is remembering that in the end, no matter what, everything will be fine. Maxime Lagacé",
			"The day you plant the seed is not the day you eat the fruit. Paulo Coelho (Source)",
			"Men who know themselves are no longer fools. They stand on the threshold of the door of wisdom. Havelock Ellis",
			"Wisdom consists not so much in knowing what to do in the ultimate as knowing what to do next. Herbert Hoover",
			"Silence at the proper season is wisdom, and better than any speech. Plutarch",
			"The most certain sign of wisdom is cheerfulness. Michel de Montaigne",
			"It’s better to be alone than to spend time with toxic people. It’s better to do nothing than to work on something that doesn’t matter. It’s better to rest than to climb the wrong mountain. James Clear",
			"Knowledge shrinks as wisdom grows. Alfred North Whitehead",
			"If the fool would persist in his folly he would become wise. William Blake",
			"Don’t try to become, let go of who you’re not. Maxime Lagacé",
			"The truest sayings are paradoxical. Lao Tzu (Tao Te Ching)",
			"Knowledge comes, but wisdom lingers. Alfred Lord Tennyson",
			"One principle eliminates a thousand decisions. Johnny Uzan",
			"Wisdom comes by disillusionment. George Santayana",
			"He that composes himself is wiser than he that composes a book. Benjamin Franklin",
			"Never does nature say one thing and wisdom another. Juvenal",
			"Pain is the doorway to wisdom and to truth. Keith Miller",
			"Kindness is wisdom. Philip James Bailey",
			"Maturity is when you stop being surprised by anything. Wisdom is when you start again. Maxime Lagacé",
			"I prefer the folly of enthusiasm to the indifference of wisdom. Anatole France",
			"The wise are those who know the Self. Sri Sathya Sai Baba",
			"No man was ever wise by chance. Seneca",
			"Wisdom begins in wonder. Socrates",
			"One part of wisdom is knowing what you don’t need anymore and letting it go. Jane Fonda",
			"Wisdom comes from experience. Experience is often a result of lack of wisdom. Terry Pratchett",
			"For everything you have missed, you have gained something else, and for everything you gain, you lose something else. Ralph Waldo Emerson",
			"The best words of wisdom will create confusion, action, understanding and ultimately, peace. Maxime Lagacé",
			"Cleverness is like a lens with a very sharp focus. Wisdom is more like a wide-angle lens. Edward de Bono",
			"It is unwise to be too sure of one’s own wisdom. It is healthy to be reminded that the strongest might weaken and the wisest might err. Mahatma Gandhi",
			"We will be known forever by the tracks we leave. Dakota (Native American saying)",
			"Has fortune dealt you some bad cards. Then let wisdom make you a good gamester. Francis Quarles",
			"The function of wisdom is to discriminate between good and evil. Marcus Tullius Cicero",
			"Genius unrefined resembles a flash of lightning, but wisdom is like the sun. Franz Grillparzer",
			"Real wisdom is not the knowledge of everything, but the knowledge of which things in life are necessary, which are less necessary, and which are completely unnecessary to know. Leo Tolstoy",
			"Wisdom is seeing wounds and obstacles as blessings. Maxime Lagacé",
			"Innocence dwells with wisdom, but never with ignorance. William Blake",
			"The small wisdom is like water in a glass: clear, transparent, pure. The great wisdom is like the water in the sea: dark, mysterious, impenetrable. Rabindranath Tagore",
			"A little nonsense now and then, is cherished by the wisest men. Roal Dahl",
		},
	}
}
