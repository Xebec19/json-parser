package internal

import "github.com/Xebec19/json-parser/pkg/tokens"

type Queue struct {
	token []tokens.Token
}

func NewQueue() *Queue {
	return &Queue{
		token: []tokens.Token{},
	}
}

func (q *Queue) Add(val tokens.Token) {

	q.token = append(q.token, val)
}

func (q *Queue) Top() (tokens.Token, error) {

	if len(q.token) == 0 {
		return "", ErrNoTokenFound
	}

	return q.token[len(q.token)-1], nil
}

func (q *Queue) Pop() (tokens.Token, error) {

	if len(q.token) == 0 {
		return "", ErrNoTokenFound
	}

	last := q.token[len(q.token)-1]

	q.token = q.token[:len(q.token)-1]

	return last, nil
}
