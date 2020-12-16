package main

import (
	"fmt"
	"io/ioutil"
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
	rules, _, nearbyTickets := parseInput("2020\\Day 16\\day16_input")
	scanningErrorRate(rules, nearbyTickets)
}

func scanningErrorRate(rules []Rule, nearbyTickets []Ticket) (errorRate int) {
	for _, ticket := range nearbyTickets {
		invalidValues := ticket.invalidValues(rules)
		for _, val := range invalidValues {
			errorRate += val
		}
	}
	fmt.Printf("Nearby ticket scanning error rate: %d\n", errorRate)
	return errorRate
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
