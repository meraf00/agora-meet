export const openMediaDevices = async (constraints: MediaStreamConstraints) => {
	return await navigator.mediaDevices.getUserMedia(constraints);
};

export const getConnectedDevices = async (type: MediaDeviceKind) => {
	const devices = await navigator.mediaDevices.enumerateDevices();
	return devices.filter((device) => device.kind === type);
};

export const onConnectedDevicesChange = (callback: (devices: MediaDeviceInfo[]) => void) => {
	navigator.mediaDevices.addEventListener('devicechange', async () => {
		const newCameraList = await getConnectedDevices('videoinput');
		callback(newCameraList);
	});
};
