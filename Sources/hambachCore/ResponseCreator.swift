import Vapor
import Foundation

public class ResponseCreator
{
    var layout: HtmlTemplate
    var navbar: HtmlTemplate
    var articleCreator: ArticleCreator
    var indexCreator: IndexCreator

    public init(layout: HtmlTemplate, navbar: HtmlTemplate, articleCreator: ArticleCreator, indexCreator: IndexCreator)
    {
        self.layout = layout
        self.navbar = navbar
        self.articleCreator = articleCreator
        self.indexCreator = indexCreator
    }

    func createResponse(content: [Content], page: String, type: String) throws -> HTTPResponse
    {
        if (page == "Index") {
            return HTTPResponse(status: .ok, body: try self.createIndexResponseBody(content: content, type: type))
        }
        return HTTPResponse(status: .ok, body: try self.createArticleResponseBody(content: content[0]))
    }

    private func createArticleResponseBody(content: Content) throws -> String
    {
        let layoutString = try self.layout.getTemplate()
        let navbarString = try self.navbar.getTemplate()
        let layout = layoutString.replacingOccurrences(of: "%navbar%", with: navbarString)
        let article = try self.articleCreator.createArticle(content: content)

        return layout.replacingOccurrences(of: "%content%", with: article)
    }

    private func createIndexResponseBody(content: [Content], type: String) throws -> String
    {
        let layoutString = try self.layout.getTemplate()
        let navbarString = try self.navbar.getTemplate()
        let layout = layoutString.replacingOccurrences(of: "%navbar%", with: navbarString)
        let article = try self.indexCreator.createIndex(content: content, type: type)

        return layout.replacingOccurrences(of: "%content%", with: article)
    }
}
