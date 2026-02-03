*dt: Task Manager*
https://github.com/shuuumz44/dt

project link: https://roadmap.sh/projects/task-tracker

Exactly what it sounds like.

# USAGE:
	-- add task
	task add "go on a walk"
	
	-- update task
	task update 14 "drive to school"

    -- delete task
	task delete 14

	-- mark to do / in progress / complete
	task mark 14 0
	
	-- list all tasks
	task list

	-- list by status
	task list 0
	task list 1
	task list 2

# FEATURES:
	@ adding (add task)

	@ updating
        -- change the name of the specified task. 

	@ deleting (delete task)

	@ marking
        -- change status of the specified task.
            0 -> todo
            1 -> doing
            2 -> done

	@ listing 
        -- list all tasks. can filter by status.

# TODO:
    * set up tests
    * expand add/delete to accept infinitely many arguments
    * add array length checking if JSON array length exceeds array capacity

# NOTES:
    * Marshal functions are for buffers. NewEncoder functions are not. If you want to read JSON,
    manipulate it, and send it back, use marshal.
    * Any time you process a buffer, you need to know the amount of bytes actually read into it. This means 
    that every time you use that buffer for something after you pass data into it, you need to specify
    the range of characters in the buffer that actually contain the data that you care about, by using a 
    slice operator. Typical usage looks like passing in "buffer[:amntBytes]"
    * JSON files might have to end in a newline, but I'm unsure. Additionally parsing JSON ruins
    its indentation (but there's functions to fix that)
    * The json package's Indent method only works on bytes.Buffer types. Which in turn only writes to a io.Reader type. This means that in order to use these you would have had to parse the JSON file in a completely different way (with the io.reader/writer methods, and possibly the json.encoder/decoder objects).
