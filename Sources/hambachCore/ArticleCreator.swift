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
        var name = "André Karrlein"
        if (content.creator == 2) {
            name = "Philipp Niedermeyer"
        }

        var article = articleLayoutString.replacingOccurrences(of: "%title%", with: content.title)
        article = article.replacingOccurrences(of: "%user%", with: name)
        article = article.replacingOccurrences(of: "%date%", with: content.date)
        article = article.replacingOccurrences(of: "%article%", with: content.article)

        return article
    }
}
