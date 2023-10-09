function showMark()
	local api = vim.api

	local bnr = vim.fn.bufnr('%')
	local ns_id = api.nvim_create_namespace('demo')

	local line_num = 5
	local col_num = 5

	local opts = {
		-- end_line = 10,
		id = 1,
		virt_text = {{"demo", "IncSearch"}},
		-- virt_text_pos = 'eol',
		-- virt_text_pos = 'overlay',
		-- virt_text_pos = 'right_align',
		virt_text_pos = 'inline',
		-- virt_text_win_col = 20,
		hl_mode = 'combine',
		strict = false,
	}

	local mark_id = api.nvim_buf_set_extmark(bnr, ns_id, line_num, col_num, opts)
end
