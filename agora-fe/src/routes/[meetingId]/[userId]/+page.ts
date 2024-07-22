export const load = ({ params }) => {
	return {
		meetingId: params.meetingId,
		userId: params.userId
	};
};
