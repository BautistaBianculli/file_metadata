package usecases

func (fileUseCase *fileUseCases) DeleteOne(fileName *string) error {
	return fileUseCase.fileRepository.DeleteOneFile(fileName)
}
