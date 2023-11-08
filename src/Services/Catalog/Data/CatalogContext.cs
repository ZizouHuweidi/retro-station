using Catalog.Data.Interfaces;
using Catalog.Entities;
using Microsoft.Extensions.Configuration;
using MongoDB.Driver;

namespace Catalog.Data {
public class CatalogContext : ICatalogContext {
  public CatalogContext(IConfiguration configuration) {
    var client = new MongoClient(
        configuration.GetValue<string>("DatabaseSettings:ConnectionString"));
    var database = client.GetDatabase(
        configuration.GetValue<string>("DatabaseSettings:DatabaseName"));

    Games = database.GetCollection<Game>(
        configuration.GetValue<string>("DatabaseSettings:CollectionName"));
    CatalogContextSeed.SeedData(Games);
  }

  public IMongoCollection<Game> Games { get; }
}
}
