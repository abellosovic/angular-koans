describe('AboutViewSpec', function() {
    it('should show "Hello, Angular" with angular template', function() {
        browser.get('http://localhost:8000/public/first_view.html');
        expect(element(by.binding('"Hello, Angular!"')).getText()).toEqual("Hello, Angular!");
    });
});
