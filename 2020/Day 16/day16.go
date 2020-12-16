package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Interval struct {
	min, max int
}

func (interval Interval) contains(n int) bool {
	return interval.min <= n && n <= interval.max
}

type Rule struct {
	name      string
	intervals []Interval
	index     int
}

func (rule Rule) allows(n int) bool {
	for _, interval := range rule.intervals {
		if interval.contains(n) {
			return true
		}
	}
	return false
}

type Ticket []int

func (ticket Ticket) invalidValues(rules []Rule) (result []int) {
	for _, num := range ticket {
		ok := false
		for _, rule := range rules {
			if rule.allows(num) {
				ok = true
				break
			}
		}
		if !ok {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	rules, myTicket, nearbyTickets := parseInput("2020\\Day 16\\day16_input")
	slightlyValidTickets, _ := discardCompletelyInvalidTickets(rules, nearbyTickets)

	allSlightlyValidTickets := append(slightlyValidTickets, myTicket)
	rulesPerIndex := matchRulesToIndices(rules, allSlightlyValidTickets)

	product := 1
	for i := 0; i < len(rules); i++ {
		rule := rulesPerIndex[i]
		if hasPrefix(rule.name, "departure") {
			product *= myTicket[i]
		}
	}
	fmt.Printf("The product of the fields on my ticket starting with \"contains\" is %d.\n", product)
}

func discardCompletelyInvalidTickets(rules []Rule, nearbyTickets []Ticket) (slightlyValidTickets []Ticket, errorRate int) {
	for _, ticket := range nearbyTickets {
		invalidValues := ticket.invalidValues(rules)
		if len(invalidValues) == 0 {
			slightlyValidTickets = append(slightlyValidTickets, ticket)
		}
		for _, val := range invalidValues {
			errorRate += val
		}
	}
	fmt.Printf("%d tickets out of %d are not completely invalid. Nearby ticket scanning error rate: %d\n",
		len(slightlyValidTickets), len(nearbyTickets), errorRate)
	return slightlyValidTickets, errorRate
}

func matchRulesToIndices(rules []Rule, tickets []Ticket) (rulesPerIndex map[int]Rule) {
	possibleRulesPerIndex := possibleRulesPerIndex(rules, tickets)

	rulesPerIndex = make(map[int]Rule)

	for len(rulesPerIndex) < len(rules) {
		ticketLength := len(tickets[0])
		for i := 0; i < ticketLength; i++ {
			if rulesPerIndex[i].name != "" {
				continue
			}
			possible := possibleRulesPerIndex[i]
			if len(possible) == 1 {
				ruleIndex := possible[0]
				rule := rules[ruleIndex]
				rulesPerIndex[i] = rule
				rule.index = i
				for j := 0; j < ticketLength; j++ {
					if i != j {
						possibleRulesPerIndex[j] = removeElement(possibleRulesPerIndex[j], ruleIndex)
					}
				}
				break
			} else if len(possible) == 0 {
				log.Panicf("No possible rules for index %d!", i)
			}
		}
	}
	return rulesPerIndex
}

func possibleRulesPerIndex(rules []Rule, tickets []Ticket) (possibleRulesPerIndex map[int][]int) {
	possibleRulesPerIndex = make(map[int][]int)
	for i := 0; i < len(tickets[0]); i++ {
		for j, rule := range rules {
			ok := true
			for _, ticket := range tickets {
				if !rule.allows(ticket[i]) {
					ok = false
					break
				}
			}
			if ok {
				possibleRulesPerIndex[i] = append(possibleRulesPerIndex[i], j)
			}
		}
	}
	return possibleRulesPerIndex
}

func parseInput(path string) (rules []Rule, myTicket Ticket, nearbyTickets []Ticket) {
	// for diversion, let's parse this without regexp
	contents, _ := ioutil.ReadFile(path)
	sections := strings.Split(string(contents), "\r\n\r\n")
	rulesSection, myTicketSection, nearbyTicketsSection := sections[0], sections[1], sections[2]
	ruleStrings := strings.Split(rulesSection, "\r\n")
	for _, ruleString := range ruleStrings {
		rules = append(rules, parseRule(ruleString))
	}
	myTicketString := strings.Split(myTicketSection, "\r\n")[1]
	myTicket = parseTicket(myTicketString)
	nearbyTicketStrings := strings.Split(nearbyTicketsSection, "\r\n")[1:]
	for _, nearbyTicketString := range nearbyTicketStrings {
		nearbyTickets = append(nearbyTickets, parseTicket(nearbyTicketString))
	}
	return rules, myTicket, nearbyTickets
}

func parseTicket(ticketString string) (ticket Ticket) {
	numStrings := strings.Split(ticketString, ",")
	for _, numString := range numStrings {
		num, _ := strconv.Atoi(numString)
		ticket = append(ticket, num)
	}
	return ticket
}

func parseRule(ruleString string) (rule Rule) {
	parts := strings.Split(ruleString, ": ")
	name, intervalPart := parts[0], parts[1]
	rule.name = name

	intervalStrings := strings.Split(intervalPart, " or ")
	for _, intervalString := range intervalStrings {
		nums := strings.Split(intervalString, "-")
		min, _ := strconv.Atoi(nums[0])
		max, _ := strconv.Atoi(nums[1])
		interval := Interval{min: min, max: max}
		rule.intervals = append(rule.intervals, interval)
	}

	return rule
}

func removeElement(slice []int, n int) []int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == n {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func hasPrefix(s, prefix string) bool {
	for i := 0; i < len(s); i++ {
		if s[:i+1] == prefix {
			return true
		}
	}
	return false
}
