<script lang="ts">
	import Video from '../video.svelte';
	import { PhoneIcon } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import { onConnectedDevicesChange } from '$lib/media';
	import {
		acceptOffer,
		addIceCandidate,
		initializeConnectionAndSendOffer,
		setRemoteDescription
	} from '$lib/rtc';

	export let data;

	let { meetingId, userId } = data;

	onMount(async () => {
		onConnectedDevicesChange(console.log);

		const socket = new WebSocket(`ws://localhost:3000/ws/${meetingId}`);

		let peerConnections: Record<string, RTCPeerConnection> = {};

		const remoteVideo = document.querySelector('#remoteVideo') as HTMLVideoElement;
		const localVideo = document.querySelector('#localVideo') as HTMLVideoElement;

		localVideo.autoplay = true;
		remoteVideo.autoplay = true;

		const localStream = await navigator.mediaDevices.getUserMedia({ video: true, audio: true });

		localVideo.srcObject = localStream;

		socket.onopen = async (e: any) => {
			console.log('Connected to', e.currentTarget.url);
			socket.send(JSON.stringify({ type: 'user:join', userId }));
		};

		socket.onmessage = async (event) => {
			console.log('Message:', event.data);

			const message = JSON.parse(event.data);

			if (message.type === 'user:join') {
				const pc = await initializeConnectionAndSendOffer(socket, userId, localStream, remoteVideo);
				peerConnections[message.userId] = pc;
			} else if (message.type === 'rtc:offer') {
				const pc = await acceptOffer(
					JSON.parse(message.payload),
					socket,
					userId,
					localStream,
					remoteVideo
				);
				peerConnections[message.userId] = pc;
			} else if (message.type === 'rtc:answer') {
				setRemoteDescription(JSON.parse(message.payload), peerConnections[message.userId]);
			} else if (message.type === 'rtc:icecandidate') {
				addIceCandidate(JSON.parse(message.payload), peerConnections[message.userId]);
			} else if (message.type === 'user:leave') {
				peerConnections[message.userId].close();
				delete peerConnections[message.userId];
			}

			console.log(peerConnections);
		};

		socket.onclose = (event) => {
			console.log('Closed:', event.code, event.reason);
			socket.send(JSON.stringify({ type: 'user:leave', userId }));
		};

		socket.onerror = (error) => {
			console.error('Error:', error);
		};
	});
</script>

<div class="p-10 h-screen overflow-hidden flex flex-col relative">
	<div class="grid grid-cols-4">
		<div class="col-span-3">
			<Video id="remoteVideo" />
		</div>
		<div class="flex flex-col gap-5 px-5 overflow-auto">
			<Video id="localVideo" />
		</div>
	</div>

	<div
		class="absolute flex justify-between py-2 px-3 items-center hover:opacity-100 opacity-0 transition-opacity duration-300
		left-1/2 -translate-x-1/2 bottom-5 w-1/3 shadow-lg rounded-full"
	>
		<button
			class="bg-red-500 hover:bg-red-700 text-white p-2 rounded-full trasnsition-colors duration-300"
		>
			<PhoneIcon class="bg-transparent" />
		</button>
	</div>
</div>
