type Stack struct {
	par []string
}

func isValid(s string) bool {
	var p Stack
	if len(s) < 2{
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			p.par = append(p.par, string(s[i]))
		}
		if len(p.par) == 0 {
    		return false
		}
		if s[i] == ')' || s[i] == '}' || s[i] == ']' {
			top := p.pop()
			if s[i] == ')' && top != "(" {
				return false
			}
			if s[i] == '}' && top != "{" {
				return false
			}
			if s[i] == ']' && top != "[" {
				return false
			}
		}

	}
	return len(p.par) == 0
}

func (s *Stack) pop() string {
	l := len(s.par) - 1
	par := s.par[l]
	s.par = s.par[:l]
	return par
}