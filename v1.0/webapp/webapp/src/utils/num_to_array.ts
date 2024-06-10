export function intToArray(i: number): Uint8Array {
    return Uint8Array.of(
        (i & 0xff000000) >> 24,
        (i & 0x0ff00000) >> 20,
        (i & 0x00ff0000) >> 16,
        (i & 0x000ff000) >> 12,
        (i & 0x0000ff00) >> 8,
        (i & 0x00000ff0) >> 4,
        (i & 0x000000ff) >> 0,

    )
}

export function intToArrayv2(i: number): Uint8Array {
    return Uint8Array.of(
        (i & 0xff000000) >> 24,
        (i & 0x00ff0000) >> 16,
        (i & 0x0000ff00) >> 8,
        (i & 0x000000ff) >> 0,
    )
}


export function Uint8ArrayToNum(arr : Uint8Array) {
    var length = arr.length;

    let buffer = Buffer.from(arr);
    var result = buffer.readUIntBE(0, length);

    return result;
}