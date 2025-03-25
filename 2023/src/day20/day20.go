package main

import (
	"2023/src/utils"
	"fmt"
	"strings"
)

type Signal struct {
	pulse               bool
	origin, destination string
}

type Module interface {
	next(Signal) []Signal
	getName() string
	getDestinations() []string
}

type FlipFlop struct {
	name         string
	status       bool
	destinations []string
}

type Conjunction struct {
	name         string
	recent       map[string]bool
	destinations []string
}

type Broadcast struct {
	name        string
	destination []string
}

func (f *FlipFlop) next(signal Signal) []Signal {
	output := make([]Signal, 0)
	if !signal.pulse {
		f.status = !f.status
		for _, dest := range f.destinations {
			output = append(output, Signal{pulse: f.status, origin: f.name, destination: dest})
		}
	}
	return output
}

func (f *FlipFlop) getName() string {
	return f.name
}

func (f *FlipFlop) getDestinations() []string {
	return f.destinations
}

func (c *Conjunction) next(signal Signal) []Signal {
	output := make([]Signal, 0)
	c.recent[signal.origin] = signal.pulse
	pulse := true
	for _, val := range c.recent {
		pulse = pulse && val
	}
	for _, dest := range c.destinations {
		output = append(output, Signal{pulse: !pulse, origin: c.name, destination: dest})
	}
	return output
}

func (c *Conjunction) getName() string {
	return c.name
}

func (c *Conjunction) getDestinations() []string {
	return c.destinations
}

func (b *Broadcast) next(signal Signal) []Signal {
	output := make([]Signal, 0)
	for _, dest := range b.destination {
		output = append(output, Signal{pulse: signal.pulse, origin: b.name, destination: dest})
	}
	return output
}

func (b *Broadcast) getName() string {
	return b.name
}

func (b *Broadcast) getDestinations() []string {
	return b.destination
}

func (s *Signal) String() string {
	return fmt.Sprintf("%s -%t-> %s", s.origin, s.pulse, s.destination)
}

func processing(filename string) map[string]Module {
	modules := make(map[string]Module)
	for _, line := range strings.Split(utils.ReadFile(filename), "\n") {
		mod, dest, _ := strings.Cut(line, " -> ")
		destinations := strings.Split(dest, ", ")
		if strings.HasPrefix(mod, "%") {
			modules[mod[1:]] = &FlipFlop{mod[1:], false, destinations}
		} else if strings.HasPrefix(mod, "&") {
			recent := make(map[string]bool)
			modules[mod[1:]] = &Conjunction{mod[1:], recent, destinations}
		} else {
			modules[mod] = &Broadcast{mod, destinations}
		}
	}
	// Conjunctions
	for _, module := range modules {
		if con, ok := module.(*Conjunction); ok {
			for _, input := range modules {
				if utils.Contains(input.getDestinations(), con.getName()) {
					con.recent[input.getName()] = false
				}
			}
		}
	}
	return modules
}

func part1(filename string) int {
	modules := processing(filename)
	count := map[bool]int{true: 0, false: 0}
	for i := 0; i < 1000; i++ {
		queue := utils.NewQueue[Signal](0)
		queue.Put(Signal{pulse: false, origin: "button", destination: "broadcaster"})
		count[false]++
		for !queue.Empty() {
			signal := queue.Get()
			if _, ok := modules[signal.destination]; !ok {
				continue
			}
			for _, sig := range modules[signal.destination].next(signal) {
				queue.Put(sig)
				count[sig.pulse]++
			}
		}
	}
	return count[true] * count[false]
}

func part2(filename string) int {
	modules := processing(filename)
	cycles, seen := make(map[string]int), make(map[string]int)
	var output string
	var push int
	for _, module := range modules {
		if utils.Contains(module.getDestinations(), "rx") {
			output = module.getName()
		}
	}
	for _, module := range modules {
		if utils.Contains(module.getDestinations(), output) {
			seen[module.getName()] = 0
		}
	}
	for {
		push++
		queue := utils.NewQueue[Signal](0)
		queue.Put(Signal{pulse: false, origin: "button", destination: "broadcaster"})
		for !queue.Empty() {
			signal := queue.Get()
			if _, ok := modules[signal.destination]; !ok {
				continue
			}
			if modules[signal.destination].getName() == output && signal.pulse {
				seen[signal.origin]++
				if _, ok := cycles[signal.origin]; !ok {
					cycles[signal.origin] = push
				}
				seenAll := true
				for _, count := range seen {
					if count == 0 {
						seenAll = false
						break
					}
				}
				if seenAll {
					x := 1
					for _, cycle := range cycles {
						x = x * cycle / utils.Gcd(x, cycle)
					}
					return x
				}
			}
			for _, sig := range modules[signal.destination].next(signal) {
				queue.Put(sig)
			}
		}
	}
}

func main() {
	filename := "day20/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
