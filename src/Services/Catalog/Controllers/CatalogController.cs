using Catalog.Entities;
using Catalog.Repositories.Interfaces;
using DnsClient.Internal;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Net;
using System.Threading.Tasks;

namespace Catalog.API.Controllers {
[ApiController]
[Route("api/v1/[controller]")]
public class CatalogController : ControllerBase {
  private readonly IGameRepository _repository;
  private readonly ILogger<CatalogController> _logger;

  public CatalogController(IGameRepository repository,
                           ILogger<CatalogController> logger) {
    _repository =
        repository ?? throw new ArgumentNullException(nameof(repository));
    _logger = logger ?? throw new ArgumentNullException(nameof(logger));
  }

  [HttpGet]
  [ProducesResponseType(typeof(IEnumerable<Game>), (int)HttpStatusCode.OK)]
  public async Task<ActionResult<IEnumerable<Game>>> GetGames() {
    var games = await _repository.GetGames();
    return Ok(games);
  }

  [HttpGet("{id:length(24)}", Name = "GetGame")]
  [ProducesResponseType((int)HttpStatusCode.NotFound)]
  [ProducesResponseType(typeof(Game), (int)HttpStatusCode.OK)]
  public async Task<ActionResult<Game>> GetGameById(string id) {
    var game = await _repository.GetGame(id);

    if (game == null) {
      _logger.LogError($"Game with id: {id}, not found.");
      return NotFound();
    }

    return Ok(game);
  }

  [Route("[action]/{category}", Name = "GetGameByCategory")]
  [HttpGet]
  [ProducesResponseType(typeof(IEnumerable<Game>), (int)HttpStatusCode.OK)]
  public async Task<ActionResult<IEnumerable<Game>>>
  GetGameByCategory(string category) {
    var games = await _repository.GetGameByCategory(category);
    return Ok(games);
  }

  [Route("[action]/{name}", Name = "GetGameByName")]
  [HttpGet]
  [ProducesResponseType((int)HttpStatusCode.NotFound)]
  [ProducesResponseType(typeof(IEnumerable<Game>), (int)HttpStatusCode.OK)]
  public async Task<ActionResult<IEnumerable<Game>>>
  GetGameByName(string name) {
    var items = await _repository.GetGameByName(name);
    if (items == null) {
      _logger.LogError($"Games with name: {name} not found.");
      return NotFound();
    }
    return Ok(items);
  }

  [HttpPost]
  [ProducesResponseType(typeof(Game), (int)HttpStatusCode.OK)]
  public async Task<ActionResult<Game>> CreateGame([FromBody] Game game) {
    await _repository.CreateGame(game);

    return CreatedAtRoute("GetGame", new { id = game.Id }, game);
  }

  [HttpPut]
  [ProducesResponseType(typeof(Game), (int)HttpStatusCode.OK)]
  public async Task<IActionResult> UpdateGame([FromBody] Game game) {
    return Ok(await _repository.UpdateGame(game));
  }

  [HttpDelete("{id:length(24)}", Name = "DeleteGame")]
  [ProducesResponseType(typeof(Game), (int)HttpStatusCode.OK)]
  public async Task<IActionResult> DeleteGameById(string id) {
    return Ok(await _repository.DeleteGame(id));
  }
}
}
