package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Rule interface {
	matches(s string, ruleRegistry RuleRegistry) bool
}

type RuleRegistry map[int]Rule

type BaseRule string // string to match

type SublistRule []int // list of Rule ids to match in order

type AlternativeRule []SublistRule // list of SublistRule of which to match any

func (baseRule BaseRule) matches(s string, _ RuleRegistry) bool {
	return s == string(baseRule)
}

func (sublistRule SublistRule) matches(toMatch string, ruleRegistry RuleRegistry) bool {
	if len(toMatch) == 0 || len(sublistRule) == 0 {
		return len(toMatch) == 0 && len(sublistRule) == 0
	}

	headRule := ruleRegistry[sublistRule[0]]
	tailRules := sublistRule[1:]

	if len(sublistRule) == 1 {
		return headRule.matches(toMatch, ruleRegistry)
	}

	maxHeadLength := len(toMatch) - len(tailRules)
	for i := 1; i <= maxHeadLength; i++ {
		// split toMatch after first i characters
		head, tail := toMatch[:i], toMatch[i:]
		if headRule.matches(head, ruleRegistry) && tailRules.matches(tail, ruleRegistry) {
			return true
		}
	}
	return false
}

func (alternativeRule AlternativeRule) matches(s string, ruleRegistry RuleRegistry) bool {
	for _, alt := range alternativeRule {
		if alt.matches(s, ruleRegistry) {
			return true
		}
	}
	return false
}

func main() {
	ruleRegistry, messages := parseInput("2020\\Day 19\\day19_input")
	fmt.Println("Processing messages without recursive rules.")
	countCompletelyMatching(messages, 0, ruleRegistry)
	maxLength := maxLengthOfStrings(messages)
	ruleRegistry[8] = buildSelfRecursiveRule(SublistRule{42}, 1, maxLength)
	ruleRegistry[11] = buildSelfRecursiveRule(SublistRule{42, 31}, 1, maxLength)
	fmt.Println("\n\nProcessing messages with recursive rules.")
	countCompletelyMatching(messages, 0, ruleRegistry)
}

func buildSelfRecursiveRule(baseRule SublistRule, recursiveIndex, maxLength int) (result AlternativeRule) {
	alternative := baseRule
	for len(alternative) <= maxLength {
		result = append(result, alternative)
		front := baseRule[:recursiveIndex]
		back := baseRule[recursiveIndex:]
		alternative = tripleAppend(front, alternative, back)
	}
	return result
}

func countCompletelyMatching(messages []string, ruleId int, ruleRegistry RuleRegistry) int {
	rule := ruleRegistry[0]

	// sync setup
	matchChannel := make(chan bool)
	numOfMessages := len(messages)

	for _, message := range messages {
		msg := message
		go testMessage(rule, msg, ruleRegistry, matchChannel)
	}

	// sum over channel entries until channel is closed
	sum := 0
	leftToProcess := numOfMessages
	for b := range matchChannel {
		if b {
			sum++
		}
		leftToProcess--
		fmt.Printf("\r%d of %d messages left to process. Found %d matches for rule %d.", leftToProcess, numOfMessages, sum, ruleId)
		if leftToProcess == 0 {
			close(matchChannel)
			break
		}
	}
	return sum
}

func testMessage(rule Rule, msg string, ruleRegistry RuleRegistry, matchChannel chan<- bool) {
	if rule.matches(msg, ruleRegistry) {
		matchChannel <- true
	} else {
		matchChannel <- false
	}
}

func parseInput(path string) (ruleRegistry RuleRegistry, messages []string) {
	ruleRegistry = make(RuleRegistry)

	contents, _ := ioutil.ReadFile(path)
	parts := strings.Split(string(contents), "\r\n\r\n")
	rulePart, messagePart := parts[0], parts[1]

	messages = strings.Split(messagePart, "\r\n")

	ruleLines := strings.Split(rulePart, "\r\n")
	ruleRegistry = parseRules(ruleLines)

	return ruleRegistry, messages
}

func parseRules(lines []string) (ruleRegistry RuleRegistry) {
	ruleRegistry = make(RuleRegistry)

	reInt := regexp.MustCompile("\\d+")
	reLetter := regexp.MustCompile("[A-z]")

	for _, line := range lines {
		id, _ := strconv.Atoi(reInt.FindString(line))
		rest := strings.Split(line, ":")[1]

		// check if BaseRule
		letter := reLetter.FindString(rest)
		if letter != "" {
			ruleRegistry[id] = BaseRule(letter)
			continue
		}

		// case: AlternativeRule
		alternativeRule := AlternativeRule{}

		alternatives := strings.Split(rest, "|")
		// loop over alternatives
		for _, alternative := range alternatives {
			sublistRule := SublistRule{}
			// loop over ids in alternative
			subRuleIdStrings := reInt.FindAllString(alternative, -1)
			for _, subRuleIdString := range subRuleIdStrings {
				subRuleId, _ := strconv.Atoi(subRuleIdString)
				sublistRule = append(sublistRule, subRuleId)
			}
			alternativeRule = append(alternativeRule, sublistRule)
		}

		if len(alternativeRule) == 1 {
			ruleRegistry[id] = alternativeRule[0]
		} else {
			ruleRegistry[id] = alternativeRule
		}
	}
	return ruleRegistry
}

func maxLengthOfStrings(strings []string) int {
	max := 0
	for _, s := range strings {
		if len(s) > max {
			max = len(s)
		}
	}
	return max
}

func tripleAppend(front, middle, back SublistRule) SublistRule {
	return append(front, append(middle, back...)...)
}
