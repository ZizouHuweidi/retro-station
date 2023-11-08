using Catalog.Data.Interfaces;
using Catalog.Entities;
using Catalog.Repositories.Interfaces;
using MongoDB.Driver;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Catalog.Repositories {
public class GameRepository : IGameRepository {
  private readonly ICatalogContext _context;

  public GameRepository(ICatalogContext context) {
    _context = context ?? throw new ArgumentNullException(nameof(context));
  }

  public async Task<IEnumerable<Game>> GetGames() {
    return await _context.Games.Find(p => true).ToListAsync();
  }

  public async Task<Game> GetGame(string id) {
    return await _context.Games.Find(p => p.Id == id).FirstOrDefaultAsync();
  }

  public async Task<IEnumerable<Game>> GetGameByName(string name) {
    FilterDefinition<Game> filter =
        Builders<Game>.Filter.ElemMatch(p => p.Name, name);

    return await _context.Games.Find(filter).ToListAsync();
  }

  public async Task<IEnumerable<Game>> GetGameByCategory(string categoryName) {
    FilterDefinition<Game> filter =
        Builders<Game>.Filter.Eq(p => p.Category, categoryName);

    return await _context.Games.Find(filter).ToListAsync();
  }

  public async Task CreateGame(Game game) {
    await _context.Games.InsertOneAsync(game);
  }

  public async Task<bool> UpdateGame(Game game) {
    var updateResult = await _context.Games.ReplaceOneAsync(
        filter: g => g.Id == game.Id, replacement: game);

    return updateResult.IsAcknowledged && updateResult.ModifiedCount > 0;
  }

  public async Task<bool> DeleteGame(string id) {
    FilterDefinition<Game> filter = Builders<Game>.Filter.Eq(p => p.Id, id);

    DeleteResult deleteResult = await _context.Games.DeleteOneAsync(filter);

    return deleteResult.IsAcknowledged && deleteResult.DeletedCount > 0;
  }
}
}
