import Vapor
import FluentMySQL
import Foundation

public class ArticleCreator
{
    var articleLayout: HtmlTemplate

    init(articleLayout: HtmlTemplate)
    {
        self.articleLayout = articleLayout
    }

    func createArticle(content: Content) throws -> String
    {
        let article = try self.replacePlaceholder(content: content)

        return article
    }

    private func replacePlaceholder(content: Content) throws -> String
    {
        let articleLayoutString = try self.articleLayout.getTemplate()
        let firstname = "André"
        let lastname = "Karrlein"

        var article = articleLayoutString.replacingOccurrences(of: "%title%", with: content.title)
        article = article.replacingOccurrences(of: "%user%", with: firstname + " " + lastname)
        article = article.replacingOccurrences(of: "%date%", with: content.date)
        article = article.replacingOccurrences(of: "%article%", with: content.article)

        return article
    }
}
