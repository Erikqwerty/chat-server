package chatservice

import "context"

func (s *service) DeleteChat(ctx context.Context, id int64) error {
	return s.chatRepository.DeleteChat(ctx, int(id))
}
