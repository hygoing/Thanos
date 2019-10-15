package DataStruct

type ItemStack struct {
	items []interface{}
}

func (s *ItemStack) init() {
	s.items = make([]interface{}, 0)
}

func (s *ItemStack) push(item interface{}) {
	items := make([]interface{}, 0)
	s.items = append(items, item, s.items)
}

func (s *ItemStack) pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[0]
	s.items = s.items[1:]
	return item
}

func (s *ItemStack) peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[0]
	return item
}
