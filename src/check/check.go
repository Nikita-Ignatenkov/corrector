package check

import (
	"context"
	"corrector/models"
	"fmt"
	"github.com/kortschak/hunspell"
	"regexp"
	"strings"
	"time"
)

const upsertCheckQuery = `
	INSERT INTO checks (created_ad) 
    VALUES ($1) 
	RETURNING id
`
const upsertBadItemsQuery = `
	INSERT INTO bad_items (name, item_id, check_id) 
    VALUES ($1, $2, $3) 
`

type Check struct {
	Id int
}

func insertCheck(db PgxDB) (int, error) {
	var id int
	if err := db.QueryRow(context.Background(), upsertCheckQuery, time.Now()).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

const updateCheckQuery = `
	UPDATE checks
	SET reviewed = $1
	WHERE id = $2
`

func updateCheck(db PgxDB, checkId int, reviewed int) error {
	fmt.Println("OPA")
	fmt.Println(checkId)
	var id int
	if err := db.QueryRow(context.Background(), updateCheckQuery, reviewed, checkId).Scan(&id); err != nil {
		fmt.Println(err)
	}

	return nil
}

func Run(db PgxDB) error {
	checkId, err := insertCheck(db)
	if err != nil {
		return err
	}

	//items, err := parser.ParseCategory(687)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(checkId)

	counter := 0

	var items [2]models.Product

	goodItem := models.Product{ID: 56, Name: "Рамка-вкладыш, учим английский язык \\\"Рыбка\\\""}
	badItem := models.Product{ID: 665, Name: "Пазл молый в рамке хуй \\\"Подъемный кран\\\", 20 элементов"}
	items[0] = goodItem
	items[1] = badItem
	fmt.Println(items)
	dictionary, err := initDictionary("/home/nikita/Projects/gitlab.sima-land.ru/nikita/showcase/src/dictionary/ru_RU.aff", "/home/nikita/Projects/gitlab.sima-land.ru/nikita/showcase/src/dictionary/ru_RU.dic")
	if err != nil {
		return err
	}

	for _, item := range items {
		words := splitWords(item.Name)
		for _, word := range words {
			result := dictionary.IsCorrect(word)

			if result == false {
				insertBadItem(db, checkId, item)
				break
			}
		}
		counter++
	}

	db.QueryRow(context.Background(), updateCheckQuery, counter, checkId)
	updateCheck(db, checkId, counter)

	return nil
}

func insertBadItem(db PgxDB, checkId int, item models.Product) {
	db.QueryRow(context.Background(), upsertBadItemsQuery, item.Name, item.ID, checkId)

	fmt.Println(item)
}

func splitWords(s string) []string {
	re := regexp.MustCompile(`[.,()\\/"]+`)
	cleanedString := re.ReplaceAllString(s, " ")
	return strings.Fields(cleanedString)
}

func initDictionary(pathToAff, pathToDic string) (*hunspell.Spell, error) {
	dictionary, err := hunspell.NewSpellPaths(pathToAff, pathToDic) // Загружаем словарь
	if err != nil {
		return nil, err
	}
	return dictionary, nil
}
