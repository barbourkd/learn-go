package dictionary

import "testing"

func TestSearch(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}

	t.Run("known word", func(t *testing.T) {
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("Bilbo")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "Kevin"
		definiton := "This is me!"

		err := dictionary.Add(word, definiton)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definiton)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "Computer"
		definition := "A thing that makes beeps and boops"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "beep boop")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}

}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
