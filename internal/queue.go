package internal

type Queue struct {
	token []Token
}

func NewQueue() *Queue {
	return &Queue{
		token: []Token{},
	}
}

func (q *Queue) Add(val Token) {

	q.token = append(q.token, val)
}

func (q *Queue) Top() (Token, error) {

	if len(q.token) == 0 {
		return "", ErrNoTokenFound
	}

	return q.token[len(q.token)-1], nil
}

func (q *Queue) Pop() (Token, error) {

	if len(q.token) == 0 {
		return "", ErrNoTokenFound
	}

	last := q.token[len(q.token)-1]

	q.token = q.token[:len(q.token)-1]

	return last, nil
}
