DIRS := A B C D E

new: .check-no # create dir and files for new contest. NO env var required
	for dir in $(DIRS); do \
		mkdir -p $(NO)/$$dir; \
		cp ./main.go $(NO)/$$dir/; \
	done
	cd $(NO)

.Phony: new

.check-no:
ifndef NO
	$(error env var NO is not set)
endif
