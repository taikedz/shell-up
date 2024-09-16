const std = @import("std");
const Writer = std.io.Writer; // FIXME - this is probably wrong

const syntax_macro = @import("syntax_macro.zig");

pub struct MacroSettings {
    .writer:Writer,
    .interpolate = true,
}

pub fn transpile(alloc:Allocator, file_name: []const u8, macro_settings: MacroSettings) {
    // register the file as seen, resolving canonical absolute path
    // open the file, start reading lines
    // if regular line, write out
    // if found file macro, `transpile(new_file, .{macro_settings.writer, true})`
    // if found syntax macro, pass line into callback, and write out
}

