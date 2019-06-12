package jsarray

// Len is part of sort.Interface.
func (s *Sorter) Len() int {
	return len(s.array)
}

// Swap is part of sort.Interface.
func (s *Sorter) Swap(i, j int) {
	s.array[i], s.array[j] = s.array[j], s.array[i]
}

// Less is part of sort.Interface.
func (s *Sorter) Less(i, j int) bool {
	return s.LessFunc(s.array[i], s.array[j])
}
