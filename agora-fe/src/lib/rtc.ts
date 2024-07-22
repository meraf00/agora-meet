import type { SignalingChannel } from './type';

const configuration = { iceServers: [{ urls: 'stun:stun.l.google.com:19302' }] };

export async function initializeConnectionAndSendOffer(
	signalingChannel: SignalingChannel,
	id: string,
	localStream: MediaStream
) {
	const peerConnection = new RTCPeerConnection(configuration);

	localStream.getTracks().forEach((track) => {
		peerConnection.addTrack(track, localStream);
	});

	peerConnection.onicecandidate = (event) => {
		if (event.candidate) {
			signalingChannel.send(
				JSON.stringify({
					userId: id,
					type: 'rtc:icecandidate',
					payload: JSON.stringify(event.candidate)
				})
			);
		}
	};

	logConnectionStateChange(peerConnection);

	const offer = await peerConnection.createOffer();
	await peerConnection.setLocalDescription(offer);

	signalingChannel.send(
		JSON.stringify({ userId: id, type: 'rtc:offer', payload: JSON.stringify(offer) })
	);

	return peerConnection;
}

export async function acceptOffer(
	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	offer: any,
	signalingChannel: SignalingChannel,
	id: string,
	localStream: MediaStream
) {
	const peerConnection = new RTCPeerConnection(configuration);

	localStream.getTracks().forEach((track) => {
		peerConnection.addTrack(track, localStream);
	});

	peerConnection.onicecandidate = (event) => {
		if (event.candidate) {
			signalingChannel.send(
				JSON.stringify({
					userId: id,
					type: 'rtc:icecandidate',
					payload: JSON.stringify(event.candidate)
				})
			);
		}
	};

	logConnectionStateChange(peerConnection);

	peerConnection.setRemoteDescription(new RTCSessionDescription(offer));

	const answer = await peerConnection.createAnswer();
	await peerConnection.setLocalDescription(answer);

	signalingChannel.send(
		JSON.stringify({ userId: id, type: 'rtc:answer', payload: JSON.stringify(answer) })
	);

	return peerConnection;
}

export async function setRemoteDescription(
	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	answer: any,
	peerConnection: RTCPeerConnection
) {
	const remoteDesc = new RTCSessionDescription(answer);
	await peerConnection.setRemoteDescription(remoteDesc);
}

export async function addIceCandidate(
	iceCandidate: RTCIceCandidate,
	peerConnection: RTCPeerConnection
) {
	if (!iceCandidate) return;
	if (!peerConnection) return;
	await peerConnection.addIceCandidate(iceCandidate);
}

export function logConnectionStateChange(peerConnection: RTCPeerConnection) {
	peerConnection.addEventListener('connectionstatechange', () => {
		console.log(peerConnection.connectionState);
	});
}
