import Vapor
import Foundation

public class ResponseCreator
{
    var layout: HtmlTemplate
    var navbar: HtmlTemplate
    var articleCreator: ArticleCreator

    public init(layout: HtmlTemplate, navbar: HtmlTemplate, articleCreator: ArticleCreator)
    {
        self.layout = layout
        self.navbar = navbar
        self.articleCreator = articleCreator
    }

    func createResponse(content: Content) throws -> HTTPResponse
    {
        let responseBody = try self.createResponseBody(content: content)
        return HTTPResponse(status: .ok, body: responseBody)
    }

    private func createResponseBody(content: Content) throws -> String
    {
        let layoutString = try self.layout.getTemplate()
        let navbarString = try self.navbar.getTemplate()
        let layout = layoutString.replacingOccurrences(of: "%navbar%", with: navbarString)
        let article = try self.articleCreator.createArticle(content: content)

        return layout.replacingOccurrences(of: "%content%", with: article)
    }
}
