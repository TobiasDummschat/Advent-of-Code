package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

type Food struct {
	ingredients, notedAllergens StringSet
}

type StringSet map[string]struct{}

func (set StringSet) Add(element string) {
	var member struct{}
	set[element] = member
}

func (set StringSet) AddAll(elements []string) {
	for _, element := range elements {
		set.Add(element)
	}
}

func (set StringSet) ToSlice() (result []string) {
	for entry := range set {
		result = append(result, entry)
	}
	return result
}

func intersection(sets ...StringSet) StringSet {
	if len(sets) == 0 {
		return make(StringSet)
	}

	intersection := make(StringSet)
	for entry := range sets[0] {
		intersection.Add(entry)
	}

	for _, set := range sets[1:] {
		for key := range intersection {
			_, ok := set[key]
			if !ok {
				delete(intersection, key)
			}
		}
	}

	return intersection
}

func union(sets ...StringSet) StringSet {
	if len(sets) == 0 {
		return make(StringSet)
	}
	union := sets[0]
	for _, set := range sets[1:] {
		for key := range set {
			union.Add(key)
		}
	}
	return union
}

func difference(minuend, subtrahend StringSet) StringSet {
	diff := make(StringSet)
	for entry := range minuend {
		diff.Add(entry)
	}
	for entry := range subtrahend {
		delete(diff, entry)
	}
	return diff
}

func main() {
	foods, allAllergens, allIngredients := parseInput("2020\\Day 21\\day21_input")

	safeIngredients, sharedIngredientsPerAllergen := safeAndUnsafeIngredients(foods, allAllergens, allIngredients)
	count := countOccurrences(safeIngredients, foods)
	fmt.Printf("%d out of %d ingredients are definitely safe. These appear a total of %d times in the %d foods.\n",
		len(safeIngredients), len(allIngredients), count, len(foods))

	ingredientPerAllergen := identifyUnsafeIngredients(allAllergens, sharedIngredientsPerAllergen)
	allergenSlice := allAllergens.ToSlice()
	sort.Strings(allergenSlice)
	fmt.Println("Allergens by ingredient:")
	for _, allergen := range allergenSlice {
		fmt.Printf("%s,", allergen)
	}
	fmt.Println()
	for _, allergen := range allergenSlice {
		fmt.Printf("%s,", ingredientPerAllergen[allergen])
	}
}

func safeAndUnsafeIngredients(foods []Food, allAllergens, allIngredients StringSet) (safeIngredients StringSet, sharedIngredientsPerAllergen map[string]StringSet) {
	unsafeIngredients := make(StringSet)
	sharedIngredientsPerAllergen = make(map[string]StringSet)

	for allergen := range allAllergens {
		var ingredientsOfFoodsContainingAllergen []StringSet
		for _, food := range foods {
			_, containsAllergen := food.notedAllergens[allergen]
			if containsAllergen {
				ingredientsOfFoodsContainingAllergen = append(ingredientsOfFoodsContainingAllergen, food.ingredients)
			}
		}
		sharedIngredients := intersection(ingredientsOfFoodsContainingAllergen...)
		sharedIngredientsPerAllergen[allergen] = sharedIngredients
		unsafeIngredients = union(unsafeIngredients, sharedIngredients)
	}

	safeIngredients = difference(allIngredients, unsafeIngredients)
	return safeIngredients, sharedIngredientsPerAllergen
}

// returns a map mapping allergens to the ingredient containing them
func identifyUnsafeIngredients(allAllergens StringSet, sharedIngredientsPerAllergen map[string]StringSet) map[string]string {
	// copy allAllergens
	unidentifiedAllergens := make(StringSet)
	for allergen := range allAllergens {
		unidentifiedAllergens.Add(allergen)
	}

	ingredientForAllergen := make(map[string]string)
	for len(unidentifiedAllergens) > 0 {
		for allergen := range unidentifiedAllergens {
			possibleIngredients := sharedIngredientsPerAllergen[allergen]
			if len(possibleIngredients) == 1 {
				for ingredient := range possibleIngredients { // only one option
					ingredientForAllergen[allergen] = ingredient
				}
				delete(unidentifiedAllergens, allergen)
				for otherAllergen := range unidentifiedAllergens {
					delete(sharedIngredientsPerAllergen[otherAllergen], ingredientForAllergen[allergen])
				}
			}
		}
	}
	return ingredientForAllergen
}

func countOccurrences(ingredients StringSet, foods []Food) int {
	count := 0
	for _, food := range foods {
		count += len(intersection(food.ingredients, ingredients))
	}
	return count
}

func parseInput(path string) (foods []Food, allAllergens, allIngredients StringSet) {
	allAllergens = make(StringSet)
	allIngredients = make(StringSet)

	contents, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(contents), "\r\n")

	reNotedAllergens := regexp.MustCompile(" \\(contains (\\w+(?:, \\w+)*)\\)")

	for _, line := range lines {
		// parse noted allergens
		allergensSingleString := reNotedAllergens.FindStringSubmatch(line)[1]
		allergenStrings := strings.Split(allergensSingleString, ", ")
		allergens := make(StringSet)
		allergens.AddAll(allergenStrings)
		allAllergens.AddAll(allergenStrings)

		// parse ingredients
		ingredientsSingleString := reNotedAllergens.ReplaceAllString(line, "")
		ingredientStrings := strings.Split(ingredientsSingleString, " ")
		ingredients := make(StringSet)
		ingredients.AddAll(ingredientStrings)
		allIngredients.AddAll(ingredientStrings)

		newFood := Food{ingredients: ingredients, notedAllergens: allergens}
		foods = append(foods, newFood)
	}
	return foods, allAllergens, allIngredients
}
