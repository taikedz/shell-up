const std = @import("std");

pub struct AppArgs {
    .filename: []u8,
    .outbin: []u8,
}

pub fn parse_app_args(allo) [][]u8 {
    const args = std.process.argsAlloc(allo);
    defer _ = std.process.argsFree(args);

    // do we have an arg parsing library as standard in zig ?
}
