*Go CLI Task Manager*

Exactly what it sounds like.

**USAGE:**
	-- add task
	task add "go on a walk"
	
	-- update task
	task update 14 "drive to school"

    -- delete task
	task delete 14

	-- mark in progress / complete
	task mark 0 14
	
	-- list all tasks
	task list

	-- list by status
	task list todo
	task list doing
	task list done

*remember: CLI arguments are 0 indexed.

**FEATURES:**
	@ adding
		-- unique key generation
		-- keys in order, 1-indexed
		-- numbers reorder when a value earlier
		than the last is completed
        -- structure is a linked list
        (you want all ids to self correct in a linked list. don't overcomplicate the obvious way to do this.)

	@ updating

	@ deleting
        -- correct indexes after list mutation

	@ marking
        -- status listed as an integer:
            0 -> todo
            1 -> doing
            2 -> done

	@ listing

to do:
    - abstract away argument length checking
    - actually add JSON I/O & parsing
