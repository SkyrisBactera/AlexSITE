var dialog = document.querySelector('dialog');
if (!dialog.showModal) {
    dialogPolyfill.registerDialog(dialog);
}
dialog.showModal();
dialog.querySelector('.close').addEventListener('click', function() {
  dialog.close();
});
