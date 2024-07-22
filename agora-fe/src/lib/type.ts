export interface SignalingChannel {
	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	send: (arg0: any) => void;
	addEventListener: (
		arg0: string,
		arg1: (arg0: { answer: RTCSessionDescriptionInit }) => void
	) => void;
}
