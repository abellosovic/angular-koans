describe('AboutModule', function() {
    it('should create module call name "gemStore" in file GemStoreModule.js', function() {
        expect(angular.module('gemStore')).not.toBe(null);
    });
});
