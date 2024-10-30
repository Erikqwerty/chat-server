package chatservice

import "context"

func (s *service) DeleteChat(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTX := s.chatRepository.DeleteChat(ctx, int(id))
		if errTX != nil {
			return errTX
		}

		if err := s.writeLog(ctx, actionTypeDeleteChat); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
