setup:
	@echo "Installing Foreman"
	@gem install foreman

run:
	foreman start -f procfile_dev