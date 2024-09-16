const std = @import("std");
const printd = std.debug.print;

const files = @import("file_load.zig");
const args = @import("arguments.zig");

const GPA = std.heap.GeneralPurposeAllocator(.{}){};

pub fn main() !void {
    defer _ = GPA.deinit();
    const alloc = GPA.allocator();

    const app_args = args.parse_app_args(alloc);
    defer _ = alloc.free(app_args);

    // Call the transpiler now
    files.transpile(alloc, app_Args.target, { .interpolate = true, .writer = filewriter })
}

