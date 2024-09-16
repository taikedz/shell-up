pub fn function_sig(allocator, line: []const u8, writer:Writer) !void {
    // turn `function_name(a *b !c ? d)` into
        // function_name() {
        //   local a="${1:-}";
        //   shift || { echo "Internal : failed to provide argument 'a'; exit 201 ; }" >&2
        //
        //   declare -n b
        //   shift || { echo "Internal : failed to provide argument 'b'; exit 201 ; }" >&2
        //
        //   c="${1:-}"
        //   shift || { echo "Internal : failed to provide argument 'b'; exit 201 ; }" >&2
        //
        //   local d="${1:-}"
        //   shift || :
    // then write the result to the given writer
}

