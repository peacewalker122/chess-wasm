// src/wasm.ts
let go: any;
let isReady = false;

export async function initWasm(): Promise<void> {
	if (isReady) return;

	go = new (window as any).Go();
	const wasm = await WebAssembly.instantiateStreaming(
		fetch("build/main.wasm"),
		go.importObject,
	);
	await go.run(wasm.instance);
	isReady = true;
}

export function createBoard(isWhiteFirst = true) {
	try {
		// Call the Go WASM function to build the board
		const result = (window as any).buildBoard(isWhiteFirst);
		console.log("Go WASM board result:", result);

		// Convert the Go board representation to our TypeScript GameState
		// return convertToGameState(result, isWhiteFirst);
	} catch (error) {
		console.error("Error creating board:", error);
		// Fallback to client-side initialization if WASM fails
		// return initializeChessGame();
	}
}
